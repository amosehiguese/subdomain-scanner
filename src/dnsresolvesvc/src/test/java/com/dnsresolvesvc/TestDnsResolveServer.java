package com.dnsresolvesvc;

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

import org.junit.After;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import com.dnsresolve.ResolveDnsServiceGrpc;
import com.dnsresolve.ResolveDnsServiceGrpc.ResolveDnsServiceStub;
import com.dnsresolve.Resolvedns.ResolveDnsRequest;
import com.dnsresolve.Resolvedns.ResolveDnsResponse;


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

        final List<ResolveDnsResponse> result = new ArrayList<ResolveDnsResponse>();
        final CountDownLatch latch = new CountDownLatch(1);
        StreamObserver<ResolveDnsResponse> respObserver = 
            new StreamObserver<ResolveDnsResponse>() {
                @Override
                public void onNext(ResolveDnsResponse value) {
                    result.add(value);              
                }

                @Override
                public void onError(Throwable t) {
                    fail();                    
                }

                @Override
                public void onCompleted() {
                    latch.countDown();
                }
            };

        ResolveDnsServiceStub stub = ResolveDnsServiceGrpc.newStub(channel);
        stub.resolveDns(req, respObserver);
        try {
            assertTrue(latch.await(1, TimeUnit.SECONDS));
        } catch (InterruptedException e) {
            System.out.println(e);
        }

        assertEquals(expected.size(), result.size());
    }

    @Test
    public void testUnknownHostsFailsToResolve() {
        ResolveDnsRequest req = ResolveDnsRequest.newBuilder()
            .addHosts("lentarex")
            .addHosts("vmrw")
            .build();
        
        final List<ResolveDnsResponse> result = new ArrayList<ResolveDnsResponse>();
        final CountDownLatch latch = new CountDownLatch(1);
        StreamObserver<ResolveDnsResponse> respObserver = 
            new StreamObserver<ResolveDnsResponse>() {
                @Override
                public void onNext(ResolveDnsResponse value) {
                    result.add(value);              
                }

                @Override
                public void onError(Throwable t) {
                    fail();                    
                }

                @Override
                public void onCompleted() {
                    latch.countDown();
                }
            };

        ResolveDnsServiceStub stub = ResolveDnsServiceGrpc.newStub(channel);
        stub.resolveDns(req, respObserver);
        try {
            assertTrue(latch.await(1, TimeUnit.SECONDS));
        } catch (InterruptedException e) {
            System.out.println(e);
        }

        assertEquals(0, result.size());
    }
}