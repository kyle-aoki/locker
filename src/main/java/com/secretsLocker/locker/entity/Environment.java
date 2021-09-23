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

    @OneToMany(cascade = CascadeType.ALL)
    public List<Secret> secrets;

    public Environment() {

    }

    public Environment(String name) {
        this.name = name;
    }
}
