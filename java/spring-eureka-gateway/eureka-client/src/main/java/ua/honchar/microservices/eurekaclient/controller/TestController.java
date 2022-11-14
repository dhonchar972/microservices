package ua.honchar.microservices.eurekaclient.controller;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TestController { // swagger: http://localhost:port/swagger-ui/index.html

    @Value("${eureka.instance.instance-id")
    private String testId;

    @GetMapping("/test-id") // gateway: http://localhost:8765/eureka-client/test-id
    public String getTestId() {
        return testId;
    }

    @GetMapping("/test") // gateway: http://localhost:8765/eureka-client/test
    public String test() {
        return "test";
    }

}
