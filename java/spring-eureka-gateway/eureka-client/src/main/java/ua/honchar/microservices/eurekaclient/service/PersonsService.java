package ua.honchar.microservices.eurekaclient.service;

import ua.honchar.microservices.eurekaclient.model.Person;

import java.util.Optional;

public interface PersonsService {

    long save(Person person);

    Optional<Person> findById(long id);

    Iterable<Person> findAll();

    long count();

    Person update(Person person);

    void delete(Person person);

}
