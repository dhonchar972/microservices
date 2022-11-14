package ua.honchar.microservices.eurekaclient.controller;

import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.*;
import ua.honchar.microservices.eurekaclient.model.Person;
import ua.honchar.microservices.eurekaclient.service.PersonsServiceImpl;

import java.util.Optional;

@RestController
@Slf4j
@AllArgsConstructor
public class TestController { // swagger: http://localhost:port/swagger-ui/index.html

    private final PersonsServiceImpl personsService;

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

    @GetMapping("/persons/{id}")
    public Optional<Person> findById(@PathVariable long id) {
        return personsService.findById(id);
    }

    @GetMapping("/persons")
    public Iterable<Person> findAll() {
        return personsService.findAll();
    }

    @GetMapping("/persons/count")
    public long count() {
        return personsService.count();
    }

    @PostMapping
    public long save(@RequestBody Person person) {
        return personsService.save(person);
    }

    @PutMapping("/person/update")
    public void update(@RequestBody Person person) {
        personsService.update(person);
    }

    @DeleteMapping
    public void delete(@RequestBody Person person) {
        personsService.delete(person);
    }

}
