package com.secretsLocker.locker.dto;

import java.util.List;

public class GetAllSecretsDto {
    public String repoName;
    public String envName;
    public List<String> secretNames;
}
