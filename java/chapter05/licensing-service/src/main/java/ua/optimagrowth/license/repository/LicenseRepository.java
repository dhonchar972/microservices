package ua.optimagrowth.license.repository;

import java.util.List;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import ua.optimagrowth.license.model.License;

/**
 * @author Dmytro Honchar <dmytro.honchar972@gmail.com>
 * Date: 11/9/2022
 */
@Repository
public interface LicenseRepository extends CrudRepository<License, String> {

    public List<License> findByOrganizationId(String organizationId);

    public License findByOrganizationIdAndLicenseId(String organizationId, String licenseId);
}
