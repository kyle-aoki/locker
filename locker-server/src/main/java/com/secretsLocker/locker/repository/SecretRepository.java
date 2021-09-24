package com.secretsLocker.locker.repository;

import com.secretsLocker.locker.entity.Secret;
import org.springframework.data.jpa.repository.JpaRepository;

public interface SecretRepository extends JpaRepository<Secret, Long> {
}
