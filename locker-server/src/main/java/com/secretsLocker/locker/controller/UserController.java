package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.ChangePasswordDto;
import com.secretsLocker.locker.dto.CreateUserDto;
import com.secretsLocker.locker.dto.DeleteUserDto;
import com.secretsLocker.locker.dto.LogInDto;
import com.secretsLocker.locker.response.MessageResponse;
import com.secretsLocker.locker.response.Response;
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
    public Response logIn(@RequestBody LogInDto logInDto) {
        logger.info("Received payload " + logInDto);
        String sessionTokenKey = userService.logIn(logInDto);
        return new MessageResponse("UR_AUTH", "Successfully logged in.");
    }

    @PostMapping("/public/auth")
    public void userIsAuthentic(
            @RequestHeader("username") String username,
            @RequestHeader("sessiontoken") String sessionToken
    ) {
        userService.authenticate(username, sessionToken);
    }

    @PostMapping("/create")
    public Response createUser(
            @RequestHeader("username") String username,
            @RequestBody CreateUserDto createUserDto
    ) {
        logger.info("Received payload " + createUserDto);
        userService.createUser(username, createUserDto);
        return new MessageResponse("UC200", "User created.");
    }

    @PostMapping("/delete")
    public Response deleteUser(@RequestBody DeleteUserDto deleteUserDto) {
        logger.info("Received payload " + deleteUserDto);
        userService.deleteUser(deleteUserDto);
        return new MessageResponse("UD200", "User deleted.");
    }

    @PostMapping("/change-password")
    public Response changePassword(
            @RequestHeader("username") String username,
            @RequestBody ChangePasswordDto changePasswordDto
    ) {
        userService.changePassword(username, changePasswordDto);
        return new MessageResponse("UCP200", "Password changed.");
    }
}
