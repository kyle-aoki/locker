package com.secretsLocker.locker.service;

import com.secretsLocker.locker.cryptography.LockerEncryption;
import com.secretsLocker.locker.cryptography.SessionToken;
import com.secretsLocker.locker.dto.ChangePasswordDto;
import com.secretsLocker.locker.dto.CreateUserDto;
import com.secretsLocker.locker.dto.DeleteUserDto;
import com.secretsLocker.locker.dto.LogInDto;
import com.secretsLocker.locker.entity.Role;
import com.secretsLocker.locker.entity.User;
import com.secretsLocker.locker.exception.Err;
import com.secretsLocker.locker.exception.UserUnauthorized;
import com.secretsLocker.locker.repository.UserRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Date;

@Service
public class UserService {

    Logger logger = LoggerFactory.getLogger(UserService.class);

    @Autowired
    UserRepository userRepository;

    public void authenticate(String username, String sessionTokenFromRequest) {
        User user = userRepository.findByUsername(username);

        if (user == null) {
            logger.info("User is unauthorized. Username is wrong or user does not exist.");
            throw new UserUnauthorized();
        }

        String sessionTokenFromDB = user.getSessionToken();
        if (sessionTokenFromDB == null) {
            logger.info("User is unauthorized. Session token null.");
            throw new UserUnauthorized();
        }

        boolean validSessionToken = SessionToken.timeSafeStringComparison(sessionTokenFromDB, sessionTokenFromRequest);
        if (!validSessionToken) {
            logger.info("User is unauthorized. Session token incorrect.");
            throw new UserUnauthorized();
        }

        if (user.getSessionTokenExpireDate().before(new Date())) {
            logger.info("User is unauthorized. Session token is expired.");
            throw new UserUnauthorized();
        }
    }

    public String logIn(LogInDto logInDto) {
        String encryptedPassword = LockerEncryption.encryptPassword(logInDto.password);
        User user = userRepository.findByUsernameAndPassword(logInDto.username, encryptedPassword);
        if (user == null) {
            logger.info("Received incorrect username or password.");
            throw new Err("WRONG_UN_PW", "Incorrect username or password.");
        }

        logger.info("User found. Setting a session token.");
        SessionToken sessionToken = new SessionToken();

        user.setSessionToken(sessionToken.getSessionTokenKey());
        user.setSessionTokenExpireDate(sessionToken.getExpireDate());

        userRepository.save(user);

        return sessionToken.getSessionTokenKey();
    }

    public void createUser(String username, CreateUserDto createUserDto) {
        User potentialUser = userRepository.findByUsername(createUserDto.username);
        if (potentialUser != null) throw new Err("USER_ALREADY_EXISTS", "User already exists.");

        if (createUserDto.role == Role.ADMIN) {
            throw new Err("ONE_ADMIN", "There can be only one admin.");
        }

        User userCreatingOtherUser = userRepository.findByUsername(username);
        if (userCreatingOtherUser.getRole() == Role.USER) {
            throw new Err("NOT_MANAGER", "Only managers can create users.");
        }

        String encryptedPassword = LockerEncryption.encryptPassword(createUserDto.password);
        User newUser = new User(createUserDto.username, encryptedPassword, createUserDto.role);

        userRepository.save(newUser);
        logger.info("Created new user " + newUser);
    }

    public void deleteUser(DeleteUserDto deleteUserDto) {
        User userToDelete = userRepository.findByUsername(deleteUserDto.username);
        if (userToDelete == null) throw new Err("USER_NOT_FOUND", "User not found.");
        userRepository.delete(userToDelete);
    }

    public void changePassword(String username, ChangePasswordDto changePasswordDto) {
        User user = userRepository.findByUsername(username);
        String encryptedPassword = LockerEncryption.encryptPassword(changePasswordDto.password);
        user.setPassword(encryptedPassword);
        userRepository.save(user);
    }
}
