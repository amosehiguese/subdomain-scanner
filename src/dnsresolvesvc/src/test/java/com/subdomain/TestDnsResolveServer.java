package com.subdomain;

import static io.grpc.MethodDescriptor.newBuilder;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;
import static org.junit.Assert.fail;

import java.util.List;
import java.util.ArrayList;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.inprocess.InProcessChannelBuilder;
import io.grpc.inprocess.InProcessServerBuilder;
import io.grpc.stub.StreamObserver;
import io.grpc.testing.GrpcCleanupRule;
import com.subdomain.DnsResolveService;

import org.junit.After;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import com.subdomain.ResolveDnsServiceGrpc.ResolveDnsServiceBlockingStub;
import com.subdomain.Resolvedns.ResolveDnsRequest;
import com.subdomain.Resolvedns.ResolveDnsResponse;


@RunWith(JUnit4.class)
public class TestDnsResolveServer {

    @Rule
    public final GrpcCleanupRule grpcCleanup = new GrpcCleanupRule();

    private ManagedChannel channel;

    @Before
    public void setup() throws Exception {
        String DnsResolveServer =
            InProcessServerBuilder.generateName();

        InProcessServerBuilder.forName(DnsResolveServer)
            .directExecutor()
            .addService(DnsResolveService.getResolveServiceImpl())
            .build()
            .start();

        channel = grpcCleanup.register(
            InProcessChannelBuilder.forName(DnsResolveServer)
                .directExecutor()
                .build()
        );

    }

    @After
    public void tearDown() {
        channel.shutdown();
    }

    @Test
    public void testDnsResolvesSuccessfully() {
        ResolveDnsRequest req = ResolveDnsRequest.newBuilder()
            .addHosts("lentarex.com")
            .addHosts("vmrw.com")
            .build();

        List<String> expected = new ArrayList<>();
        expected.add("lentarex.com");
        expected.add("vmrw.com");

        ResolveDnsServiceBlockingStub stub = ResolveDnsServiceGrpc.newBlockingStub(channel);
        ResolveDnsResponse result = stub.resolveDns(req);

        assertEquals(expected.size(), result.getSubdomainCount());
    }

    @Test
    public void testUnknownHostsFailsToResolve() {
        ResolveDnsRequest req = ResolveDnsRequest.newBuilder()
            .addHosts("lentarex")
            .addHosts("vmrw")
            .build();

        ResolveDnsServiceBlockingStub stub = ResolveDnsServiceGrpc.newBlockingStub(channel);
        ResolveDnsResponse result = stub.resolveDns(req);

        assertEquals(0, result.getSubdomainCount());
    }
}
