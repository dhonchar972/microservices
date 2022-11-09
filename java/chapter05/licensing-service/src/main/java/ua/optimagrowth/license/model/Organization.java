package ua.optimagrowth.license.model;

import org.springframework.hateoas.RepresentationModel;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

/**
 * @author Dmytro Honchar <dmytro.honchar972@gmail.com>
 * Date: 11/9/2022
 */
@Getter
@Setter
@ToString
public class Organization extends RepresentationModel<Organization> {

    String id;
    String name;
    String contactName;
    String contactEmail;
    String contactPhone;
}
