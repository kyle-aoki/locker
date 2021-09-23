package com.secretsLocker.locker.entity;

import javax.persistence.*;
import java.util.List;

@Entity
@Table(name = "repos")
public class Repository {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    long id;

    @Column(unique = true)
    String name;

    @ManyToOne
    User owner;

    @OneToMany(cascade = CascadeType.ALL)
    List<Environment> environments;

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
