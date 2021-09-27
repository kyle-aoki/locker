package com.secretsLocker.locker.response;

import com.secretsLocker.locker.entity.Repository;

import java.util.List;

public class RepoResponse {
    public static class RepoCreated extends Response {
        public String code = "RR001";
        public String message = "Repository created.";
        public List<Repository> list;
    }
    public static class RepoList extends Response {
        public String code = "RR002";
        public List<String> list;
        public RepoList(List<String> list) {
            this.list = list;
        }
    }
    public static class RepoUpdated extends Response {
        public String code = "RR003";
        public String message = "Repository name updated.";
    }
}
