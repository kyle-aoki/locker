package com.secretsLocker.locker.exception;

import com.secretsLocker.locker.exception.handler.ExceptionResponse;
import com.secretsLocker.locker.exception.handler.ServerException;

public class EnvironmentException {
    public static class EnvironmentNameTaken extends ServerException {
        public EnvironmentNameTaken() {
            super(new ExceptionResponse("EE001","Environment name taken."));
        }
    }

}
