package com.secretsLocker.locker.repository;

import com.secretsLocker.locker.entity.Repository;
import org.springframework.data.jpa.repository.JpaRepository;

@org.springframework.stereotype.Repository
public interface RepoRepository extends JpaRepository<Repository, Long> {
    Repository findByName(String name);
}
