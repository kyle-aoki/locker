package com.secretsLocker.locker.exception;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

@ControllerAdvice
public class ExceptionAdvice {

    @ResponseBody
    @ExceptionHandler(value = UserUnauthorized.class)
    public ResponseEntity<ExceptionResponse> authException(Err err) {
        ExceptionResponse er = new ExceptionResponse(err);
        return new ResponseEntity<>(er, HttpStatus.UNAUTHORIZED);
    }

    @ResponseBody
    @ExceptionHandler(value = Err.class)
    public ExceptionResponse exception(Err err) {
        return new ExceptionResponse(err);
    }
}
