package ua.optimagrowth.license.config;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

import lombok.Getter;
import lombok.Setter;

/**
 * @author Dmytro Honchar <dmytro.honchar972@gmail.com>
 * Date: 11/9/2022
 */
@Configuration
@ConfigurationProperties(prefix = "example")
@Getter
@Setter
public class ServiceConfig {

    private String property;
}
