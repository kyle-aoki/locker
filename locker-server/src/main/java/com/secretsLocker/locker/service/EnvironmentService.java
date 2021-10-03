package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.*;
import com.secretsLocker.locker.dto.diff.MissingRequest;
import com.secretsLocker.locker.dto.path.RepoEnvPath;
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

    public Environment findByNameOrThrow(Repository repo, String envName) {
        Environment env = null;
        for (Environment e : repo.environments) {
            if (e.name.equals(envName)) env = e;
        }
        if (env == null) throw new Err("ENV_NOT_FOUND", "Environment not found.");
        return env;
    }

    public void create(RepoEnvPath repoEnvPath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvPath.repoName);

        List<Environment> envs = repo.getEnvironments();

        String newEnvName = repoEnvPath.envName;

        for (Environment env : envs) {
            if (env.name.equals(newEnvName)) throw new Err("ENV_NAME_TAKEN", "Environment name taken.");
        }
        envs.add(new Environment(newEnvName));

        repoRepository.save(repo);
    }

    public List<KeyValue> get(RepoEnvPath repoEnvPath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvPath.repoName);
        Environment env = this.findByNameOrThrow(repo, repoEnvPath.envName);

        List<KeyValue> keyValues = new ArrayList<>();
        for (Secret s : env.secrets) keyValues.add(new KeyValue(s.name, s.value));

        return keyValues;
    }

    public List<String> missing(MissingRequest missingRequest) {
        Repository repo = repoService.findByNameOrThrow(missingRequest.repoName);
        Environment givenEnv = this.findByNameOrThrow(repo, missingRequest.envName);
        Environment targetEnv = this.findByNameOrThrow(repo, missingRequest.targetEnvName);

        List<Secret> givenEnvSecrets = new ArrayList<>(givenEnv.secrets);
        List<Secret> targetEnvSecrets = new ArrayList<>(targetEnv.secrets);

        targetEnvSecrets.removeAll(givenEnvSecrets);

        return targetEnvSecrets.stream()
                .map(secret -> secret.name)
                .collect(Collectors.toList());
    }
}





















