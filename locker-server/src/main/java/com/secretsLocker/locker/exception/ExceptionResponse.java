package com.secretsLocker.locker.exception;

public class ExceptionResponse {
    public Boolean ok = false;
    public String code;
    public String message;

    public ExceptionResponse(Err err) {
        this.code = err.code;
        this.message = err.message;
    }
}
