package com.example.catsapi.repository;

import com.example.catsapi.model.Cat;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CatRepository extends JpaRepository<Cat, String> {

}
