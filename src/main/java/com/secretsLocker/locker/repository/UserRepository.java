package com.secretsLocker.locker.repository;

import com.secretsLocker.locker.cryptography.SessionToken;
import com.secretsLocker.locker.entity.Role;
import com.secretsLocker.locker.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface UserRepository extends JpaRepository<User, Integer> {
    User findByUsername(String username);
    User findByUsernameAndPassword(String username, String password);

//    @Query(value = """
//            SELECT *
//            FROM users
//            WHERE username = :username
//            AND session_token = :session_token
//    """, nativeQuery = true)
//    User getUserByUsernameAndSessionToken(
//            @Param("username") String username,
//            @Param("session_token") String sessionToken,
//    );

    @Query(value = """
            
            """)
    User findByUsernameAndSessionToken(String username, SessionToken sessionToken);

    @Query(value = """
            SELECT count(*) FROM users WHERE role = 'ADMIN';
            """, nativeQuery = true)
    Long countAdmins();

    @Query(value = """
            SELECT count(*) FROM users WHERE password = 'p4ssw0rd' and role = 'ADMIN';
            """, nativeQuery = true)
    Long countDefaultAdmins();
    User findByPasswordAndRole(String password, Role role);
}
