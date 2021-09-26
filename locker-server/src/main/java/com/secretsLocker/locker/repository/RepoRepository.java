package com.secretsLocker.locker.repository;

import com.secretsLocker.locker.entity.Repository;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

@org.springframework.stereotype.Repository
public interface RepoRepository extends JpaRepository<Repository, Long> {
    Repository findByName(String name);

    @Query(value = """
            SELECT name FROM repos LIMIT :limit OFFSET :offset
            """, nativeQuery = true)
    List<String> findAllWithLimitAndOffset(int limit, int offset);

    @Query(value = """
            SELECT name FROM repos LIMIT :limit
            """, nativeQuery = true)
    List<String> findAllNamesWithLimit(int limit);

    @Query(value = """
            SELECT name FROM repos
            """, nativeQuery = true)
    List<String> findAllNames();


}
