import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { DocsController } from './docs.controller';
import { DocsService } from './docs.service';
import { Doc, DocSchema } from '../schemas/doc.schema';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: Doc.name, schema: DocSchema }]),
  ],
  controllers: [DocsController],
  providers: [DocsService],
})
export class DocsModule {}
