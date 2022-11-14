package ua.honchar.microservices.eurekaclient.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import javax.validation.constraints.Size;
import java.util.List;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Entity
@Table(name = "organizations")
public class Organization {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Size(min = 3, max = 15)
    @Column(nullable = false)
    private String name;

    @Column(nullable = false, length = 20)
    private Country country;

    @OneToMany(mappedBy="persons")
    @JoinColumn(name = "personsId", referencedColumnName = "id")
    private List<Person> employs;

}
