package com.secretsLocker.locker.exception;

import com.secretsLocker.locker.exception.handler.ExceptionResponse;
import com.secretsLocker.locker.exception.handler.ServerException;

public class SecretException {

    public static class SecretAlreadyExists extends ServerException {
        public SecretAlreadyExists() {
            super(new ExceptionResponse("SE001","Secret already exists in this environment."));
        }
    }
    public static class SecretDoesNotExist extends ServerException {
        public SecretDoesNotExist() {
            super(new ExceptionResponse("SE002","Secret does not exist."));
        }
    }
}
