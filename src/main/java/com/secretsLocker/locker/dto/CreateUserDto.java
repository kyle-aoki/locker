package com.secretsLocker.locker.dto;

import com.secretsLocker.locker.entity.Role;

public class CreateUserDto {
    public String username;
    public String password;
    public Role role;

    @Override
    public String toString() {
        return "CreateUserDto{" +
                "username='" + username + '\'' +
                ", password='" + password + '\'' +
                ", role=" + role +
                '}';
    }
}
