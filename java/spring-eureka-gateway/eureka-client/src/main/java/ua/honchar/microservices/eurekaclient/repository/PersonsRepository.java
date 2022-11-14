package ua.honchar.microservices.eurekaclient.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;
import ua.honchar.microservices.eurekaclient.model.Person;

@Repository // optional
public interface PersonsRepository extends CrudRepository<Person, Long> {
}
