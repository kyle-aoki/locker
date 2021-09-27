package com.secretsLocker.locker.entity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "users")
public class User {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    public long id;

    @Column(unique = true, nullable = false)
    public String username;

    @Column(nullable = false)
    public String password;

    @Column
    public String sessionToken;

    @Column
    @Temporal(TemporalType.TIMESTAMP)
    public Date sessionTokenExpireDate;

    @Enumerated(EnumType.STRING)
    @Column(name = "role")
    public Role role;

    public User() {

    }

    public User(String username, String password, Role role) {
        this.username = username;
        this.password = password;
        this.role = role;
    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getSessionToken() {
        return sessionToken;
    }

    public void setSessionToken(String sessionToken) {
        this.sessionToken = sessionToken;
    }

    public Date getSessionTokenExpireDate() {
        return sessionTokenExpireDate;
    }

    public void setSessionTokenExpireDate(Date sessionTokenExpireDate) {
        this.sessionTokenExpireDate = sessionTokenExpireDate;
    }

    public Role getRole() {
        return role;
    }

    public void setRole(Role role) {
        this.role = role;
    }
}
