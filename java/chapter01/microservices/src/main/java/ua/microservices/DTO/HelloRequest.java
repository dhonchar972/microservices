package ua.microservices.DTO;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * @author Dmytro Honchar <dmytro.honchar972@gmail.com>
 * Date: 11/7/2022
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class HelloRequest {

    private String firstName;
    private String lastName;
}
