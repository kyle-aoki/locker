package com.secretsLocker.locker.dto;

import com.secretsLocker.locker.entity.Role;

public class LogInDto {
    public String username;
    public String password;

    @Override
    public String toString() {
        return "CreateUserDto{" +
                "username='" + username + '\'' +
                ", password='" + password + '\'' +
                '}';
    }
}
