package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.cryptography.SessionToken;
import com.secretsLocker.locker.dto.ChangePasswordDto;
import com.secretsLocker.locker.dto.CreateUserDto;
import com.secretsLocker.locker.dto.DeleteUserDto;
import com.secretsLocker.locker.dto.LogInDto;
import com.secretsLocker.locker.response.UserResponse;
import com.secretsLocker.locker.service.UserService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(value = "/user")
public class UserController {

    @Autowired
    UserService userService;

    Logger logger = LoggerFactory.getLogger(UserController.class);

    @PostMapping("/public/log-in")
    public UserResponse.LoggedIn logIn(@RequestBody LogInDto logInDto) {
        logger.info("Received payload " + logInDto);
        String sessionTokenKey = userService.logIn(logInDto);
        return new UserResponse.LoggedIn(sessionTokenKey);
    }

    @PostMapping("/public/auth")
    public void userIsAuthentic(
            @RequestHeader("username") String username,
            @RequestHeader("sessiontoken") String sessionToken
    ) {
        userService.authenticate(username, sessionToken);
    }

    @PostMapping("/create")
    public UserResponse.UserCreated createUser(
            @RequestHeader("username") String username,
            @RequestBody CreateUserDto createUserDto
    ) {
        logger.info("Received payload " + createUserDto);
        userService.createUser(username, createUserDto);
        return new UserResponse.UserCreated();
    }

    @PostMapping("/delete")
    public UserResponse.UserDeleted deleteUser(@RequestBody DeleteUserDto deleteUserDto) {
        logger.info("Received payload " + deleteUserDto);
        userService.deleteUser(deleteUserDto);
        return new UserResponse.UserDeleted();
    }

    @PostMapping("/change-password")
    public UserResponse.PasswordChanged changePassword(
            @RequestHeader("username") String username,
            @RequestBody ChangePasswordDto changePasswordDto
    ) {
        userService.changePassword(username, changePasswordDto);
        return new UserResponse.PasswordChanged();
    }
}
