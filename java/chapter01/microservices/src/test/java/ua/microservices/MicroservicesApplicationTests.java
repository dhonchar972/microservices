package ua.microservices;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.jupiter.api.Assertions;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.ResponseEntity;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import org.springframework.web.client.RestTemplate;
import ua.microservices.DTO.HelloRequest;

/**
 * @author Dmytro Honchar <dmytro.honchar972@gmail.com>
 * Date: 11/7/2022
 */
@SpringBootTest
@ExtendWith(SpringExtension.class)
public class MicroservicesApplicationTests {

    private static final RestTemplate restTemplate = new RestTemplate();
    private static final SpringApplicationBuilder uws = new SpringApplicationBuilder(MicroservicesApplication.class);

    @BeforeAll
    static void init() {
        uws.run();
    }

    @Test
    public void testMethodGet_returnString() {
        ResponseEntity<String> response
                = restTemplate.getForEntity("http://localhost:8080/hello/dmytro?lastName=honchar", String.class);

        Assertions.assertEquals("{\"message\":\"Hello dmytro honchar\"}", response.getBody());
    }

    @Test
    public void testMethodPost_returnString() {
        var request = new HelloRequest("dmytro", "honchar");

        ResponseEntity<String> response
                = restTemplate.postForEntity("http://localhost:8080/hello/", request, String.class);

        Assertions.assertEquals("{\"message\":\"Hello dmytro honchar\"}", response.getBody());
    }

}
