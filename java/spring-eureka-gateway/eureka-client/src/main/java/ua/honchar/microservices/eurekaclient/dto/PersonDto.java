package ua.honchar.microservices.eurekaclient.dto;

import lombok.*;

@Getter
@Setter
@ToString
@AllArgsConstructor
public record PersonDto(String name, int age, String country) {
}
