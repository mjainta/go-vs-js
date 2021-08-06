import { MongooseModule } from '@nestjs/mongoose';
import { Test, TestingModule } from '@nestjs/testing';
import { Doc, DocSchema } from '../schemas/doc.schema';
import { DocsController } from './docs.controller';
import { DocsService } from './docs.service';

describe('DocsController', () => {
  let controller: DocsController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      imports: [MongooseModule.forFeature([{ name: Doc.name, schema: DocSchema }]),
      MongooseModule.forRoot('mongodb://localhost/test_database')],
      controllers: [DocsController],
      providers: [DocsService],
    }).compile();

    controller = module.get<DocsController>(DocsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
    expect(controller.getAllDocuments).toBeDefined();
    expect(controller.getFilteredDocuments).toBeDefined();
  });
});
