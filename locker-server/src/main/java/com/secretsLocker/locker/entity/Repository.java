package com.secretsLocker.locker.entity;

import javax.persistence.*;
import java.util.List;

@Entity
@Table(name = "repos")
public class Repository {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    public long id;

    public String name;

    @ManyToOne
    public User owner;

    @ManyToMany(fetch = FetchType.LAZY)
    public List<User> members;

    @OneToMany(fetch = FetchType.LAZY)
    public List<Environment> environments;

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public User getOwner() {
        return owner;
    }

    public void setOwner(User owner) {
        this.owner = owner;
    }

    public List<Environment> getEnvironments() {
        return environments;
    }

    public void setEnvironments(List<Environment> environments) {
        this.environments = environments;
    }
}
