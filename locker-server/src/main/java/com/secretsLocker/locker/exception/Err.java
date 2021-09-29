package com.secretsLocker.locker.exception;

public class Err extends RuntimeException {
    public String code;
    public String message;

    public Err(String code, String message) {
        this.code = code;
        this.message = message;
    }
}
