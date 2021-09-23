package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.CreateEnvironmentDto;
import com.secretsLocker.locker.entity.Environment;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.exception.EnvironmentException;
import com.secretsLocker.locker.repository.RepoRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class EnvironmentService {

    @Autowired
    RepoRepository repoRepository;

    public void create(CreateEnvironmentDto createEnvironmentDto) {
        Repository repo = repoRepository.findByName(createEnvironmentDto.repoName);

        List<Environment> envs = repo.getEnvironments();

        String newEnvName = createEnvironmentDto.envName;

        for (Environment env : envs) {
            if (env.name.equals(newEnvName)) throw new EnvironmentException.EnvironmentNameTaken();
        }
        envs.add(new Environment(newEnvName));

        repoRepository.save(repo);
    }
}
