//package ua.honchar.microservices.eurekaclient.config;
//
//import com.netflix.discovery.DiscoveryClient;
//import io.github.resilience4j.circuitbreaker.annotation.CircuitBreaker;
//import lombok.AllArgsConstructor;
//import org.springframework.cloud.client.circuitbreaker.CircuitBreakerFactory;
//import org.springframework.cloud.openfeign.FeignClient;
//import org.springframework.web.bind.annotation.GetMapping;
//import org.springframework.web.bind.annotation.RequestParam;
//import org.springframework.web.client.RestTemplate;
//
//@FeignClient(name="service-name")
//public interface Client {
//    private final DiscoveryClient discoveryClient;
//    private  final RestTemplate restTemplate;
//    private  final CircuitBreakerFactory circuitBreakerFactory;
//
//    @GetMapping("/service-name")
//    List<ServiceName> getNameFromServiceNameById(@RequestParam Long id);
//}
