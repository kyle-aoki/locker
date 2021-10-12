package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.*;
import com.secretsLocker.locker.dto.path.RepoEnvPath;
import com.secretsLocker.locker.dto.path.RepoEnvSecretPath;
import com.secretsLocker.locker.dto.path.RepoEnvSecretValuePath;
import com.secretsLocker.locker.entity.Environment;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.Secret;
import com.secretsLocker.locker.exception.Err;
import com.secretsLocker.locker.repository.EnvironmentRepository;
import com.secretsLocker.locker.repository.RepoRepository;
import com.secretsLocker.locker.repository.SecretRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class SecretService {

    @Autowired RepoRepository repoRepository;
    @Autowired EnvironmentRepository environmentRepository;
    @Autowired SecretRepository secretRepository;
    @Autowired EnvironmentService environmentService;
    @Autowired RepoService repoService;

    public Secret findByName(Environment env, String secretName) {
        Secret secret = null;
        for (Secret s : env.secrets) if (s.name.equals(secretName)) secret = s;
        return secret;
    }

    public Secret findByNameOrThrow(Environment env, String secretName) {
        Secret secret = this.findByName(env, secretName);
        if (secret == null) throw new Err("SECRET_NOT_FOUND", "Secret not found.");
        return secret;
    }

    public void create(RepoEnvSecretValuePath repoEnvSecretValuePath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvSecretValuePath.repoName);
        Environment env = environmentService.findByNameOrThrow(repo, repoEnvSecretValuePath.envName);
        Secret secret = this.findByName(env, repoEnvSecretValuePath.secretName);

        if (secret != null) throw new Err("SECRET_ALREADY_EXISTS", "Secret already exists.");

        Secret newSecret = new Secret(repoEnvSecretValuePath.secretName, repoEnvSecretValuePath.secretValue);
        env.secrets.add(newSecret);

        secretRepository.save(newSecret);
        environmentRepository.save(env);
        repoRepository.save(repo);
    }

    public void update(RepoEnvSecretValuePath repoEnvSecretValuePath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvSecretValuePath.repoName);
        Environment env = environmentService.findByNameOrThrow(repo, repoEnvSecretValuePath.envName);
        Secret secret = this.findByNameOrThrow(env, repoEnvSecretValuePath.secretName);

        secret.value = repoEnvSecretValuePath.secretValue;
        secretRepository.save(secret);
    }

    public void rename(RenameSecretDto renameSecretDto) {
        Repository repo = repoService.findByNameOrThrow(renameSecretDto.repoName);
        Environment env = environmentService.findByNameOrThrow(repo, renameSecretDto.envName);
        Secret secret = this.findByNameOrThrow(env, renameSecretDto.secretName);

        secret.name = renameSecretDto.newSecretName;
        secretRepository.save(secret);
    }

    public List<KeyValue> get(RepoEnvPath repoEnvPath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvPath.repoName);
        Environment env = environmentService.findByNameOrThrow(repo, repoEnvPath.envName);

        List<KeyValue> keyValues = new ArrayList<>();
        for (Secret s : env.secrets) keyValues.add(new KeyValue(s.name, s.value));

        return keyValues;
    }

    public String getSecret(RepoEnvSecretPath repoEnvSecretPath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvSecretPath.repoName);
        Environment env = environmentService.findByNameOrThrow(repo, repoEnvSecretPath.envName);
        Secret secret = this.findByNameOrThrow(env, repoEnvSecretPath.secretName);

        return secret.value;
    }

    public List<String> list(RepoEnvPath repoEnvPath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvPath.repoName);
        Environment env = environmentService.findByNameOrThrow(repo, repoEnvPath.envName);

        return this.buildList(env);
    }

    public List<String> buildList(Environment env) {
        List<String> secretNames = new ArrayList<>();
        for (Secret s : env.secrets) secretNames.add(s.name);

        return secretNames;
    }

    public void delete(RepoEnvSecretPath repoEnvSecretPath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvSecretPath.repoName);
        Environment env = environmentService.findByNameOrThrow(repo, repoEnvSecretPath.envName);
        Secret secret = this.findByNameOrThrow(env, repoEnvSecretPath.secretName);

        env.secrets.remove(secret);

        environmentRepository.save(env);
    }
}
