package com.secretsLocker.locker.exception.handler;

public class ExceptionResponse {
    public Boolean ok = false;
    public String code;
    public String message;

    public ExceptionResponse(String code, String message) {
        this.code = code;
        this.message = message;
    }
}
