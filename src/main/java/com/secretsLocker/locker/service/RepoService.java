package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.CreateRepoDto;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.User;
import com.secretsLocker.locker.repository.RepoRepository;
import com.secretsLocker.locker.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class RepoService {

    @Autowired
    UserRepository userRepository;

    @Autowired
    RepoRepository repoRepository;

    public void create(String username, CreateRepoDto createRepoDto) {
        User user = userRepository.findByUsername(username);

        Repository repository = new Repository();

        repository.setName(createRepoDto.repoName);
        repository.setOwner(user);

        repoRepository.save(repository);
    }

}
