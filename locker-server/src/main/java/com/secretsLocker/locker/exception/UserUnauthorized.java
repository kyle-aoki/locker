package com.secretsLocker.locker.exception;

public class UserUnauthorized extends Err {
    public UserUnauthorized() {
        super("UEAUTH", "User unauthorized.");
    }
}
