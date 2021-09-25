package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.CreateRepoDto;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.User;
import com.secretsLocker.locker.exception.RepoException;
import com.secretsLocker.locker.repository.RepoRepository;
import com.secretsLocker.locker.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class RepoService {

    @Autowired
    UserRepository userRepository;

    @Autowired
    RepoRepository repoRepository;

    public void create(String username, CreateRepoDto createRepoDto) {
        User user = userRepository.findByUsername(username);

        Repository potentialRepo = repoRepository.findByName(createRepoDto.repoName);
        if (potentialRepo != null) throw new RepoException.RepoAlreadyExists();

        Repository repository = new Repository();

        repository.setName(createRepoDto.repoName);
        repository.setOwner(user);

        repoRepository.save(repository);
    }

    public List<String> listRepos() {
        List<Repository> repos = repoRepository.findAll();
        List<String> repoNames = new ArrayList<>();
        for (Repository repo : repos) {
            repoNames.add(repo.name);
        }

        return repoNames;
    }

}
