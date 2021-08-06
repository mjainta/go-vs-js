import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { DocsModule } from './docs/docs.module';

@Module({
  imports: [
    MongooseModule.forRoot('mongodb://localhost/mydatabase'),
    DocsModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
