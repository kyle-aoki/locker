package com.secretsLocker.locker.response;

import com.secretsLocker.locker.entity.Repository;

import java.util.List;

public class RepoResponse extends Response {
    public static class RepoCreated extends RepoResponse {
        public String code = "RR001";
        public String message = "Repository created.";
        public List<Repository> list;
    }
    public static class RepoList extends RepoResponse {
        public String code = "RR002";
        public List<String> list;
        public RepoList(List<String> list) {
            this.list = list;
        }
    }
}
