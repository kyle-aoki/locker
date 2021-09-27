package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.CreateEnvironmentDto;
import com.secretsLocker.locker.dto.GetEnvDto;
import com.secretsLocker.locker.dto.RawSecret;
import com.secretsLocker.locker.entity.Environment;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.Secret;
import com.secretsLocker.locker.exception.EnvironmentException;
import com.secretsLocker.locker.exception.RepoException;
import com.secretsLocker.locker.repository.RepoRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

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
        if (env == null) throw new EnvironmentException.EnvironmentNotFound();
        return env;
    }

    public void create(CreateEnvironmentDto createEnvironmentDto) {
        Repository repo = repoService.findByName(createEnvironmentDto.repoName);

        List<Environment> envs = repo.getEnvironments();

        String newEnvName = createEnvironmentDto.envName;

        for (Environment env : envs) {
            if (env.name.equals(newEnvName)) throw new EnvironmentException.EnvironmentNameTaken();
        }
        envs.add(new Environment(newEnvName));

        repoRepository.save(repo);
    }

    public List<RawSecret> get(GetEnvDto getEnvDto) {
        Repository repo = repoService.findByName(getEnvDto.repoName);
        Environment env = this.findByName(repo, getEnvDto.envName);

        List<RawSecret> rawSecrets = new ArrayList<>();
        for (Secret s : env.secrets) rawSecrets.add(new RawSecret(s.name, s.value));

        return rawSecrets;
    }
}
