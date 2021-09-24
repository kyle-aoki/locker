package com.secretsLocker.locker.entity;

import javax.persistence.*;
import java.util.List;

@Entity
@Table(name = "environments")
public class Environment {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    public long id;

    @Column(unique = true)
    public String name;

    @ManyToMany(cascade = CascadeType.ALL)
    public List<User> members;

    @OneToMany(cascade = CascadeType.ALL)
    public List<Secret> secrets;

    public Environment() {

    }

    public Environment(String name) {
        this.name = name;
    }

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

    public List<Secret> getSecrets() {
        return secrets;
    }

    public void setSecrets(List<Secret> secrets) {
        this.secrets = secrets;
    }
}
