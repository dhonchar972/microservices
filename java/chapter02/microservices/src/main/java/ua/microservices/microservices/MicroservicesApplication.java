package ua.microservices.microservices;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

/**
 * @author Dmytro Honchar <dmytro.honchar972@gmail.com>
 * Date: 11/7/2022
 */
@SpringBootApplication
@RestController
@RequestMapping(value="hello")
@EnableEurekaClient
public class MicroservicesApplication {

	public static void main(String[] args) {
		SpringApplication.run(MicroservicesApplication.class, args);
	}

	public String helloRemoteServiceCall(String firstName, String lastName) {
		RestTemplate restTemplate = new RestTemplate();

		ResponseEntity<String> restExchange =
				restTemplate.exchange("http://logical-service-id/name/" + "{firstName}/ {lastName}",
                        HttpMethod.GET, null, String.class, firstName, lastName);

		return restExchange.getBody();
	}

	@RequestMapping(value="/{firstName}/{lastName}", method = RequestMethod.GET)
	public String hello(@PathVariable("firstName") String firstName,
						@PathVariable("lastName") String lastName) {
		return helloRemoteServiceCall(firstName, lastName);
	}

}
