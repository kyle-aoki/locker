package com.secretsLocker.locker.response;

import com.secretsLocker.locker.cryptography.SessionToken;

public class UserResponse extends Response {
    public static class UserCreated extends UserResponse {
        public String code = "UR001";
        public String message = "User Created.";
    }

    public static class UserDeleted extends UserResponse {
        public String code = "UR002";
        public String message = "User Deleted.";
    }

    public static class LoggedIn extends UserResponse {
        public String code = "UR003";
        public String message = "Successfully logged in.";
        public String sessionToken;
        public LoggedIn(String sessionTokenKey) {
            this.sessionToken = sessionTokenKey;
        }
    }

    public static class PasswordChanged extends UserResponse {
        public String code = "UR004";
        public String message = "Password Changed.";
    }
}
