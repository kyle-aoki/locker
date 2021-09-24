package com.secretsLocker.locker.exception;

import com.secretsLocker.locker.exception.handler.ExceptionResponse;
import com.secretsLocker.locker.exception.handler.ServerException;

public class RepoException {
    public static class RepoAlreadyExists extends ServerException {
        public RepoAlreadyExists() {
            super(new ExceptionResponse("RE001","Repository already exists."));
        }
    }
    public static class RepoDoesNotExist extends ServerException {
        public RepoDoesNotExist() {
            super(new ExceptionResponse("RE002","Repository does not exist."));
        }
    }
}
