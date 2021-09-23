package com.secretsLocker.locker.entity;

import javax.persistence.*;

@Entity
@Table(name = "secrets")
public class Secret {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    long id;

    String name;
    String value;
}
