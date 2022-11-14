package ua.honchar.microservices.eurekaclient.service;

import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import ua.honchar.microservices.eurekaclient.model.Person;
import ua.honchar.microservices.eurekaclient.repository.PersonsRepository;

import java.util.Optional;

@AllArgsConstructor
@Slf4j
@Service
public class PersonsServiceImpl implements PersonsService {

    private final PersonsRepository personsRepository;

    @Override
    public long save(Person person) {
        var p = personsRepository.save(person);

        log.info("Saving person");

        return p.getId();
    }

    @Override
    public Optional<Person> findById(long id) {
        return personsRepository.findById(id);
    }

    @Override
    public Iterable<Person> findAll() {
        return personsRepository.findAll();
    }

    @Override
    public long count() {
        return personsRepository.count();
    }

    @Override
    public Person update(Person person) {
        return personsRepository.findById(person.getId()).map(p -> {
            p.setName(person.getName());
            p.setAge(person.getAge());
            p.setCountry(person.getCountry());
            return p;
        }).orElseGet(() -> personsRepository.save(person));
    }

    @Override
    public void delete(Person person) {
        if (personsRepository.existsById(person.getId()))
            personsRepository.delete(person);

        log.info("Deleted person " + person.getId());
    }

}
