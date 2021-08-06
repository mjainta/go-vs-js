import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from './../src/app.module';

describe('AppController (e2e)', () => {
  let app: INestApplication;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();
  });

  it('/ (GET)', () => {
    return request(app.getHttpServer())
      .get('/')
      .expect(200)
      .expect('Hello World!');
  });

  it('/documents (GET)', () => {
    return request(app.getHttpServer())
      .get('/documents')
      .expect(200)
      .expect('[{"_id":"61097a12f064fd0007ba1c18","name":"this is my first name"},{"_id":"61097a19f064fd0007ba1c1a","name":"this is my second name"},{"_id":"61097a22f064fd0007ba1c1c","name":"and my last name"}]');
  });
});
