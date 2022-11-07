package ua.microservices.controller;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;
import ua.microservices.DTO.HelloRequest;

/**
 * @author Dmytro Honchar <dmytro.honchar972@gmail.com>
 * Date: 11/7/2022
 */
@SpringBootApplication
@RestController
@RequestMapping(value="hello")
public class HelloController {

    @GetMapping(value="/{firstName}")
    public String helloGET(
            @PathVariable("firstName") String firstName,
            @RequestParam("lastName") String lastName) {
        return String.format("{\"message\":\"Hello %s %s\"}",firstName, lastName);
    }

    @PostMapping
    public String helloPOST( @RequestBody HelloRequest request) {
        return String.format("{\"message\":\"Hello %s %s\"}",request.getFirstName(), request.getLastName());
    }
}
