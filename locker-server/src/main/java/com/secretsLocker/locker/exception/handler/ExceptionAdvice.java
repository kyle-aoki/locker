package com.secretsLocker.locker.exception.handler;

import com.secretsLocker.locker.exception.UserException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

@ControllerAdvice
public class ExceptionAdvice {

    @ResponseBody
    @ExceptionHandler(value = UserException.UserUnauthorized.class)
    public ResponseEntity<ExceptionResponse> authException(ServerException serverException) {
        return new ResponseEntity<>(serverException.exceptionResponse, HttpStatus.UNAUTHORIZED);
    }

    @ResponseBody
    @ExceptionHandler(value = ServerException.class)
    public ExceptionResponse exception(ServerException serverException) {
        return serverException.exceptionResponse;
    }
}
