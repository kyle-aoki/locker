package com.secretsLocker.locker.exception.handler;

public class ServerException extends RuntimeException {
    public ExceptionResponse exceptionResponse;
    public ServerException(ExceptionResponse exceptionResponse) {
        this.exceptionResponse = exceptionResponse;
    }
}
