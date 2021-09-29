package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.*;
import com.secretsLocker.locker.dto.diff.MissingRequest;
import com.secretsLocker.locker.entity.Environment;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.Secret;
import com.secretsLocker.locker.exception.Err;
import com.secretsLocker.locker.repository.RepoRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@Service
public class EnvironmentService {

    @Autowired
    RepoRepository repoRepository;

    @Autowired
    RepoService repoService;

    public Environment findByName(Repository repo, String envName) {
        Environment env = null;
        for (Environment e : repo.environments) {
            if (e.name.equals(envName)) env = e;
        }
        if (env == null) throw new Err("ENV_NOT_FOUND", "Environment not found.");
        return env;
    }

    public void create(CreateEnvironmentDto createEnvironmentDto) {
        Repository repo = repoService.findByNameOrThrow(createEnvironmentDto.repoName);

        List<Environment> envs = repo.getEnvironments();

        String newEnvName = createEnvironmentDto.envName;

        for (Environment env : envs) {
            if (env.name.equals(newEnvName)) throw new Err("ENV_NAME_TAKEN", "Environment name taken.");
        }
        envs.add(new Environment(newEnvName));

        repoRepository.save(repo);
    }

    public List<KeyValue> get(GetEnvDto getEnvDto) {
        Repository repo = repoService.findByNameOrThrow(getEnvDto.repoName);
        Environment env = this.findByName(repo, getEnvDto.envName);

        List<KeyValue> keyValues = new ArrayList<>();
        for (Secret s : env.secrets) keyValues.add(new KeyValue(s.name, s.value));

        return keyValues;
    }

    public List<String> missing(MissingRequest missingRequest) {
        Repository repo = repoService.findByNameOrThrow(missingRequest.repoName);
        Environment givenEnv = this.findByName(repo, missingRequest.envName);
        Environment targetEnv = this.findByName(repo, missingRequest.targetEnvName);

        List<Secret> givenEnvSecrets = new ArrayList<>(givenEnv.secrets);
        List<Secret> targetEnvSecrets = new ArrayList<>(targetEnv.secrets);

        targetEnvSecrets.removeAll(givenEnvSecrets);

        return targetEnvSecrets.stream()
                .map(secret -> secret.name)
                .collect(Collectors.toList());
    }
}





















