package com.secretsLocker.locker.exception;

import com.secretsLocker.locker.exception.handler.ExceptionResponse;
import com.secretsLocker.locker.exception.handler.ServerException;

public class UserException {
    public static class UserAlreadyExists extends ServerException {
        public UserAlreadyExists() {
            super(new ExceptionResponse("UE001","User already exists."));
        }
    }
    public static class UserNotFound extends ServerException {
        public UserNotFound() {
            super(new ExceptionResponse("UE002", "User Not Found."));
        }
    }
    public static class IncorrectUsernameOrPassword extends ServerException {
        public IncorrectUsernameOrPassword() {
            super(new ExceptionResponse("UE003", "Incorrect username or password."));
        }
    }
    public static class UserUnauthorized extends ServerException {
        public UserUnauthorized() {
            super(new ExceptionResponse("UEAUTH", "User unauthorized."));
        }
    }
    public static class CannotCreateAdditionalAdmins extends ServerException {
        public CannotCreateAdditionalAdmins() {
            super(new ExceptionResponse("UE004", "There can be only one admin."));
        }
    }
}
