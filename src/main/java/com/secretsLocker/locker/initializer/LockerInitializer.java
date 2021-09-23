package com.secretsLocker.locker.initializer;

import com.secretsLocker.locker.cryptography.LockerEncryption;
import com.secretsLocker.locker.entity.Role;
import com.secretsLocker.locker.entity.User;
import com.secretsLocker.locker.repository.UserRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.util.Date;

@Component
class LockerInitializer {

    Logger logger = LoggerFactory.getLogger(LockerInitializer.class);

    @Autowired
    UserRepository userRepository;

    private static final String DefaultAdminUsername = "admin";
    private static final String DefaultAdminPassword = "p4ssw0rd";
    private static final String DefaultAdminEncryptedPassword = LockerEncryption.encryptPassword(DefaultAdminPassword);

    @PostConstruct
    private void initializeAdminAccountIfAbsent() {
        logger.info("Checking admin status...");

        Long adminCount = userRepository.countAdmins();

        if (adminCount == 0) {
            User defaultAdmin = new User();

            defaultAdmin.setUsername(DefaultAdminUsername);
            defaultAdmin.setPassword(DefaultAdminEncryptedPassword);
            defaultAdmin.setRole(Role.ADMIN);
            defaultAdmin.setSessionToken("default");
            defaultAdmin.setSessionTokenExpireDate(new Date(System.currentTimeMillis() + 9999999L));

            userRepository.save(defaultAdmin);

            this.printAdminInitWarningMessage();

            return;
        }

        Long defaultAdminCount = userRepository.countDefaultAdmins();
        User defaultAdmin = userRepository.findByPasswordAndRole(DefaultAdminEncryptedPassword, Role.ADMIN);

        if (defaultAdmin != null) {
            this.printDefaultAdminWarningMessage();
        }
    }

    private void printAdminInitWarningMessage() {
        logger.warn("******************** WARNING ********************");
        logger.warn("******************** WARNING ********************");
        logger.warn("Default admin account initialized. Username is 'admin', password is 'p4ssw0rd'.");
        logger.warn("admin password is set to 'p4ssw0rd'. [username=admin password=p4ssw0rd]");
        logger.warn("Please change the default password into something more secure.");
        logger.warn("******************** WARNING ********************");
        logger.warn("******************** WARNING ********************");
    }

    private void printDefaultAdminWarningMessage() {
        logger.warn("******************** WARNING ********************");
        logger.warn("******************** WARNING ********************");
        logger.warn("admin password is set to 'p4ssw0rd'. [username=admin password=p4ssw0rd]");
        logger.warn("Please change the default password into something more secure.");
        logger.warn("******************** WARNING ********************");
        logger.warn("******************** WARNING ********************");
    }
}
