package ua.microservices.DTO;

import lombok.Data;

/**
 * @author Dmytro Honchar <dmytro.honchar972@gmail.com>
 * Date: 11/7/2022
 */
@Data
public class HelloRequest {

    private String firstName;
    private String lastName;
}
